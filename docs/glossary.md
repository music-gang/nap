# Glossary

The words Nap uses, in the docs, in the logs, and in the code. If another page
disagrees with this one, this page wins. Code-only terms are in
[development](development.md#terms-used-in-the-code).

## Words shared with Nomad

| Word | In Nap | In Nomad |
|---|---|---|
| **deployment** | Not used for Nap's own actions. | The rolling update Nomad runs after a register: write **Nomad deployment**. |
| **job** | A Nomad job. Nap never changes one. | The same. |
| **count** | The number of allocations a group runs. The only value Nap writes. | The same. |
| **version** | The Nap build (`v0.1.0`): write **Nap version**. | A job's version: write **job version**. |

## Groups and targets

| Term | Meaning |
|---|---|
| **target** | A group Nap serves, written `namespace/job/group`. |
| **managed group** | A group whose meta has `nap_enabled = "true"` and that has an enabled `scaling` block with `min = 0`. |
| **managed namespace** | A namespace in `-nomad-namespaces`. |
| **idle period** | How long a target gets no request before Nap puts it to sleep: `nap_idle`, else the `-default-idle` option. |

## Waking and sleeping

| Term | Meaning |
|---|---|
| **awake** | A target at count 1. |
| **asleep** | A target at count 0. |
| **wake** | Scale a target from 0 to 1 because a request arrived. |
| **sleep** | Scale a target from 1 to 0 because it stayed idle. |
| **ready** | A target whose allocations are healthy and whose ready delay has passed. |
| **ready delay** | The margin after the allocations turn healthy, so the proxy sees them: `-ready-delay`. |
| **forwardAuth** | The proxy feature Nap uses: the proxy asks Nap about each request before it forwards it. |
