issuer: https://{{ .GlobalConfig.Get "fqdn" }}/cas/oidc
fqdn: https://{{ .GlobalConfig.Get "fqdn" }}
port: 8080
client_id: {{ .Config.GetAndDecrypt "sa-cas/oidc_client_id" }}
client_secret: {{ .Config.GetAndDecrypt "sa-cas/oidc_client_secret" }}
