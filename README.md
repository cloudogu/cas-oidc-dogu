# cas-oidc-dogu

Dogu for a simple oidc client. Used for testing purposes only

### Local development

1. Start CAS locally via `.gradlew build run`
2. Start client `go run .`

### Testing CAS-OIDC via Dogu

1. Build and start dogu with `cesapp build .` and `cesapp start cas-oidc-dogu`.
2. Visit `<fqdn>/cas-oidc-dogu/login`.
3. Browser should recieve profile as JSON:

```json
{
  "attributes": {
    "cn": "admin",
    "displayName": "admin",
    "givenName": "admin",
    "groups": [
      "cesManager",
      "asd"
    ],
    "mail": "asd@asd.asd",
    "surname": "admin",
    "username": "asd"
  },
  "auth_time": 1665040091,
  "client_id": "cas-oidc-client",
  "id": "asd",
  "service": "https://192.168.56.2/cas-oidc-client/auth/callback",
  "sub": "asd"
}
```

## What is the Cloudogu EcoSystem?
The Cloudogu EcoSystem is an open platform, which lets you choose how and where your team creates great software. Each service or tool is delivered as a Dogu, a Docker container. Each Dogu can easily be integrated in your environment just by pulling it from our registry.

We have a growing number of ready-to-use Dogus, e.g. SCM-Manager, Jenkins, Nexus Repository, SonarQube, Redmine and many more. Every Dogu can be tailored to your specific needs. Take advantage of a central authentication service, a dynamic navigation, that lets you easily switch between the web UIs and a smart configuration magic, which automatically detects and responds to dependencies between Dogus.

The Cloudogu EcoSystem is open source and it runs either on-premises or in the cloud. The Cloudogu EcoSystem is developed by Cloudogu GmbH under [AGPL-3.0-only](https://spdx.org/licenses/AGPL-3.0-only.html).

## License
Copyright © 2020 - present Cloudogu GmbH
This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, version 3.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
You should have received a copy of the GNU Affero General Public License along with this program. If not, see https://www.gnu.org/licenses/.
See [LICENSE](LICENSE) for details.


---
MADE WITH :heart:&nbsp;FOR DEV ADDICTS. [Legal notice / Imprint](https://cloudogu.com/en/imprint/?mtm_campaign=ecosystem&mtm_kwd=imprint&mtm_source=github&mtm_medium=link)
