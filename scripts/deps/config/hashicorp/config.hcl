backend "file" {
    path            = "/vault/file"
}

listener "tcp" {
    address         = "vault:8200"
    tls_disable     = true
}

default_lease_ttl = "30s"
max_lease_ttl = "1m"

log_level = "Debug"

ui = true
