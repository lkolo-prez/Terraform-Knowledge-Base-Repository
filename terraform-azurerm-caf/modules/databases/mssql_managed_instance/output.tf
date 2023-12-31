
output "name" {
  value       = azurecaf_name.mssqlmi.result
  description = "SQL MI Name"
}

output "id" {
  value       = local.output.id
  description = "SQL MI Id"
}

output "location" {
  value = var.location
}

output "principal_id" {
  value       = local.output.principal_id
  description = "SQL MI Identity Principal Id"
}

output "resource_group_id" {
  value       = local.parent_id
  description = "Resource group resource id of the SQL Server managed instance."
}