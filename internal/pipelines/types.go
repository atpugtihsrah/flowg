package pipelines

import (
	"context"

	"link-society.com/flowg/internal/storage"
	"link-society.com/flowg/internal/vrl"
)

type Pipeline struct {
	Root Node
}

func (p *Pipeline) Run(
	ctx context.Context,
	manager *Manager,
	entry *storage.LogEntry,
) error {
	return p.Root.Process(ctx, manager, entry)
}

type Node interface {
	Process(ctx context.Context, manager *Manager, entry *storage.LogEntry) error
}

type TransformNode struct {
	TransformerName string
	Next            []Node
}

type SwitchNode struct {
	Branches map[string]SwitchNodeBranch
}

type SwitchNodeBranch struct {
	Condition storage.Filter
	Next      []Node
}

type RouterNode struct {
	Stream string
}

func (n *TransformNode) Process(
	ctx context.Context,
	manager *Manager,
	entry *storage.LogEntry,
) error {
	vrlScript, err := manager.GetTransformerScript(n.TransformerName)
	if err != nil {
		return err
	}

	output, err := vrl.ProcessRecord(entry.Fields, vrlScript)
	if err != nil {
		return err
	}

	for _, next := range n.Next {
		newEntry := &storage.LogEntry{
			Timestamp: entry.Timestamp,
			Fields:    output,
		}
		err := next.Process(ctx, manager, newEntry)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *SwitchNode) Process(
	ctx context.Context,
	manager *Manager,
	entry *storage.LogEntry,
) error {
	for _, branch := range n.Branches {
		if branch.Condition.Evaluate(entry) {
			for _, next := range branch.Next {
				err := next.Process(ctx, manager, entry)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (n *RouterNode) Process(
	ctx context.Context,
	manager *Manager,
	entry *storage.LogEntry,
) error {
	_, err := manager.db.Append(ctx, n.Stream, entry)
	return err
}
