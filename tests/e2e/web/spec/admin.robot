*** Settings ***
Library   SeleniumLibrary
Library   RequestsLibrary
Resource  common.resource

*** Test Cases ***
Create test role
    Open Browser                   ${BASE_URL}  ${BROWSER}
    Log as                         username=root  password=root
    Click Navbar Menu Item         id=menu:navbar.profile  id=link:navbar.profile.admin
    Wait Until Page Contains       Roles
    Click Element                  id=btn:admin.roles.create
    Wait Until Element Is Visible  id=input:admin.roles.modal.name
    Input Text                     id=input:admin.roles.modal.name  test
    Element Should Be Visible      xpath=//div[@id='field:admin.roles.modal.scopes']//div[@data-ref='container:generic.transfer-list.items-left']//div[@id='label:generic.transfer-list.read_pipelines']
    Click Element                  xpath=//div[@id='field:admin.roles.modal.scopes']//input[@id='checkbox:generic.transfer-list.read_pipelines']
    Click Element                  xpath=//div[@id='field:admin.roles.modal.scopes']//button[@id='btn:generic.transfer-list.selected-right']
    Wait Until Element Is Visible  xpath=//div[@id='field:admin.roles.modal.scopes']//div[@data-ref='container:generic.transfer-list.items-right']//div[@id='label:generic.transfer-list.read_pipelines']  timeout=5s
    Element Should Not Be Visible  xpath=//div[@id='field:admin.roles.modal.scopes']//div[@data-ref='container:generic.transfer-list.items-left']//div[@id='label:generic.transfer-list.read_pipelines']
    Click Element                  id=btn:admin.roles.modal.submit
    Wait Until Page Contains       Role created  timeout=5s
    Wait Until Element Is Visible  xpath=//div[@id='table:admin.roles']//div[@role='row'][@row-id='test']  timeout=5s
    Close Browser
