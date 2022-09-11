This is your metas folder.

Any YAML file (`.yaml` or `.yml`) in this folder (or subfolder) will be considered a meta for your expansion.

You can apply the `meta_schema.json` schema to the YAML files in this folder to help you.

A meta is simply a Set (archetype) or a counter. Here is a list of all properties a meta can have.

| Property name | Description                                              | Type               | Required |
| ------------- | -------------------------------------------------------- | ------------------ | -------- |
| `id`          | The id of this meta                                      | Integer            | ✅ Yes    |
| `type`        | The type of this meta                                    | "set" or "counter" | ✅ Yes    |
| `name`        | The name of this meta                                    | String             | ✅ Yes    |
| `alias`       | An alias for a **set** to be used on a card `sets` field | String             | ❌ No     |

If an alias is not provided to a set, use its `name` instead.
