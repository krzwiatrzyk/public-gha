resource "null_resource" "a" {}

module "c" {
    source = "../../modules/m1"
}