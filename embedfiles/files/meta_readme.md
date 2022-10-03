This is your metas folder.

Any YAML file (`.yaml` or `.yml`) in this folder (or subfolder) will be considered as a meta for your expansion.

You can apply the `meta_schema.json` schema to the YAML files in this folder to help you.

A meta is simply a Set (archetype), a counter or an alias. Here is a list of all properties a meta can have.

| Property name | Description                                              | Type                        | Required |
| ------------- | -------------------------------------------------------- | --------------------------- | -------- |
| `id`          | The id of this meta                                      | Integer                     | ✅ Yes    |
| `type`        | The type of this meta                                    | "set", "counter" or "alias" | ✅ Yes    |
| `name`        | The name of this meta                                    | String                      | ✅ Yes    |
| `alias`       | An alias for a **set** to be used on a card `sets` field | String                      | ❌ No     |

If the `alias` field is not provided to a set, use its `name` instead.

# Aliases

A meta of type `alias` has no actual function. It's only purpose is to be used on a `meta` macro as an alias to a longer text (See `macros.md.txt` on the `cards` folder).
