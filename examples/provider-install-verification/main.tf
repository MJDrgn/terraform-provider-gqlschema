terraform {
  required_providers {
    gql = {
      source = "registry.terraform.io/mjdrgn/gqlschema"
    }
  }
}

provider "gql" {}

output "queryfield" {
  value = data.gql_field.queryfield.schema
}

data "gql_field" "queryfield" {
  id = "listThings"
  type = "[Thing!]!"

  argument {
    id = "limit"
    type = "Int"
    default = "10"
  }

  argument {
    id = "offset"
    type = "Int"
    default = "0"
  }
}

output "example" {
  value = data.gql_type.example.schema
}

data "gql_type" "example" {
  id = "example"

  field {
    id = "name"
    type = "String!"
  }

  field {
    id = "Attributes"
    type = "[String!]!"
  }

  field {
    id = "Children"
    type = "[String!]!"

    argument {
      id = "limit"
      type = "Int"
      default = "10"
    }

    argument {
      id = "offset"
      type = "Int"
      default = "0"
    }

    argument {
      id = "filter"
      type = "String"

      directive {
        id = "deprecated"

        argument {
          id = "reason"
          value = "Reason goes here"
        }
      }
    }
  }

  field {
    id = "Old"
    type = "String"

    directive {
      id = "deprecated"
    }
  }

}
