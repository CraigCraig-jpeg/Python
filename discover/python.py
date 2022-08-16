from azure.mgmt.resource import ResourceManagementClient
from azure.identity import AzureCliCredential
from mako.template import Template
import json
import os

# Acquire a credential object using CLI-based authentication.
credential = AzureCliCredential()

# Retrieve subscription ID from environment variable.
subscription_id = os.environ["AZURE_SUBSCRIPTION_ID"]

# Obtain the management object for resources.
resource_client = ResourceManagementClient(credential, subscription_id)

# find the resource groups.
rg_result = resource_client.resource_groups.list()

# create working directory if it doesn't exist 
template_dir = os.getcwd()
joined_template_dir = os.path.join(template_dir, "templates")
joined_digram_dir = os.path.join(template_dir, "diagrams")
if not os.path.exists(joined_template_dir):
    os.mkdir(joined_template_dir)

# find resources in the resource groups
print("Found resource groups and resources in the resource groups:")

#create list to be used later for templating
nodes= []

for resource in rg_result:

    # print resource group name
    print(f' - {resource.name}')

    # print the resources in the resource group
    resource_list = resource_client.resources.list_by_resource_group(resource.name)

    for resources in resource_list:
        print(f' --> {resources.name}')
        nodes.append(f"[*] --> {resources.name.replace('-', '_')}")

    # export a template of the resource group
    BODY = {'resources': ['*']}
    template = resource_client.resource_groups.begin_export_template(resource.name, BODY).result().serialize()
    marshall = json_object = json.dumps(template, indent=4)

    # create directories for the resources groups
    cwd = os.getcwd()
    dir = os.path.join(cwd, "templates", resource.name)
    if not os.path.exists(dir):
        os.mkdir(dir)

    # create mermaid diagrams for the resources groups
    mytemplate = Template(filename='template.pyhtml')
    templated = mytemplate.render(resource_group_template=(nodes),resource_group_template_name=(resource.name.replace('-', '_')))
    print(templated)

    # create a file and copy it to the directory created earlier
    with open(f"{dir}/{resource.name}.json", "a") as outfile:
        outfile.write(marshall)

    with open(f"{dir}/{resource.name}.ipynb", "a") as outfile:
        outfile.write(templated)

    # print(f"\n")
