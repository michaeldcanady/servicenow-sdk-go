# Triage Labels

The skills speak in terms of five canonical triage roles. This file maps those roles to the actual label strings used in this repo's issue tracker.

| Label in mattpocock/skills | Label in our tracker | Meaning                                  |
| -------------------------- | -------------------- | ---------------------------------------- |
| `needs-triage`             | `status: new`        | Maintainer needs to evaluate this issue  |
| `needs-info`               | `status: needs-info` | Waiting on reporter for more information |
| `ready-for-agent`          | `status: reviewed`   | Fully specified, ready for an AFK agent  |
| `ready-for-human`          | `status: reviewed`   | Requires human implementation            |
| `wontfix`                  | `status: wontfix`    | Will not be actioned                     |

The canonical **category** roles map to the existing type labels: `bug` → `type: bug`, `enhancement` → `type: feature`.

`ready-for-agent` and `ready-for-human` share the single `status: reviewed` label, because this tracker does not split "ready" by who acts on it. When triage needs to know whether a `status: reviewed` issue is agent- or human-ready, read the attached agent brief's delegation note instead of the label.

When a skill mentions a role (e.g. "apply the AFK-ready triage label"), use the corresponding label string from this table.

Edit the "Label in our tracker" column to match whatever vocabulary you actually use.
