document.addEventListener('DOMContentLoaded', () => {
  const action_save = document.getElementById('action_save')
  const data_transformer_name = document.getElementById('data_transformer_name')
  const data_transformer_code = document.getElementById('data_transformer_code')

  if (data_transformer_name.value !== '') {
    history.pushState(null, '', `/web/transformers/edit/${data_transformer_name.value}/`)
  }

  action_save.addEventListener('click', () => {
    if (data_transformer_name.value === '') {
      M.toast({ html: '&#10060; Please provide a transformer name' })
      data_transformer_name.classList.add('invalid')
    } else {
      const form = document.createElement('form')
      form.setAttribute('method', 'post')
      form.setAttribute('action', window.location.href)
      form.classList.add('hide')

      const input_name = document.createElement('input')
      input_name.setAttribute('type', 'hidden')
      input_name.setAttribute('name', 'name')
      input_name.setAttribute('value', data_transformer_name.value)

      const input_code = document.createElement('input')
      input_code.setAttribute('type', 'hidden')
      input_code.setAttribute('name', 'code')
      input_code.setAttribute('value', data_transformer_code.getAttribute('code'))

      form.appendChild(input_name)
      form.appendChild(input_code)
      document.body.appendChild(form)

      form.submit()
    }
  })

  data_transformer_name.addEventListener('input', () => {
    data_transformer_name.classList.remove('invalid')
  })
})
