This is your cards folder.

Any YAML file (`.yaml` or `.yml`) in this folder (or subfolder) will be considered a card for your expansion.

You can apply the `card_schema.json` schema to the YAML files in this folder to help you.

Here is a list of all properties a card can have.

| Property name     | Description                       | Type                         | Required |
| ----------------- | --------------------------------- | ---------------------------- | -------- |
| `id`              | The code of your card             | Non-zero integer             | ✅ Yes    |
| `name`            | Card name                         | String                       | ✅ Yes    |
| `description`     | Effect box content or flavor text | String                       | ✅ Yes    |
| `card_type`       | Main type of this card            | "monster", "spell" or "trap" | ✅ Yes    |
| `sub_types`       | Sub-types of this card            | String array*                | ❌ No     |
| `ruleset`         | Ruleset this card applies         | String array*                | ❌ No     |
| `sets`            | Sets this card belongs to         | Set alias or id array        | ❌ No     |
| `atk`             | Monster attack points             | Integer (Negative is `?`)    | ❌ No     |
| `def`             | Monster defense points            | Integer (Negative is `?`)    | ❌ No     |
| `level`           | Monster level/rank/link-rating    | Integer                      | ❌ No     |
| `race`            | Monster type                      | String array*                | ❌ No     |
| `attribute`       | Monster attribute                 | String array*                | ❌ No     |
| `category`        | Search categories of this card    | String array*                | ❌ No     |
| `strings`         | Strings for effects               | String array (up to 16)      | ❌ No     |
| `pendulum_effect` | Pendulum box content              | String                       | ❌ No     |
| `link_arrows`     | Link-arrows directions            | String array*                | ❌ No     |
| `scale`           | Pendulum scale                    | Integer                      | ❌ No     |
| `alias`           | Alias card code                   | Non-zero Integer             | ❌ No     |

String arrays marked with * means they have a list of available values. You can see them below.


## Available sub-types
| Sub-type         | Sub type for                                 |
| ---------------- | -------------------------------------------- |
| `normal`         | Normal monsters                              |
| `effect`         | Effect monsters                              |
| `fusion`         | Fusion monsters                              |
| `ritual`         | Ritual spells                                |
| `spirit`         | Spirit monsters                              |
| `union`          | Union monsters                               |
| `gemini`         | Gemini monsters                              |
| `tuner`          | Tuner monsters                               |
| `synchro`        | Synchro monsters                             |
| `token`          | Tokens                                       |
| `quickplay`      | Quick-Play spells                            |
| `continuous`     | Continuous spells/traps                      |
| `equip`          | Equip spells                                 |
| `field`          | Field spells                                 |
| `counter`        | Counter traps                                |
| `flip`           | Flip monsters                                |
| `toon`           | Toon monsters                                |
| `xyz`            | XYZ monsters                                 |
| `pendulum`       | Pendulum monsters                            |
| `special_summon` | Monsters that must be special summoned first |
| `link`           | Link monsters                                |
| `skill`          | Skill cards                                  |
| `action`         | Action spells/traps                          |
| `plus`           | Plus monsters                                |
| `minor`          | Minus monsters                               |
| `armor`          | Armor monsters                               |


## Available rulesets
| Ruleset      | Ruleset for |
| ------------ | ----------- |
| `ocg`        | OCG         |
| `tcg`        | TCG         |
| `animanga`   | Anime       |
| `illegal`    | Illegal     |
| `videogame`  | Video-games |
| `custom`     | Custom      |
| `speed`      | Speed Duel  |
| `prerelease` | Pre-release |
| `rush`       | Rush Duel   |
| `legend`     | Legend      |
| `hidden`     | Hidden      |


## Available races
| Race            | Race for              |
| --------------- | --------------------- |
| `warrior`       | Warrior monster       |
| `spellcaster`   | Spellcaster monster   |
| `fairy`         | Fairy monster         |
| `fiend`         | Fiend monster         |
| `zombie`        | Zombie monster        |
| `machine`       | Machine monster       |
| `aqua`          | Aqua monster          |
| `pyro`          | Pyro monster          |
| `rock`          | Rock monster          |
| `winged_beast`  | Winged Beast monster  |
| `plant`         | Plant monster         |
| `insect`        | Insect monster        |
| `thunder`       | Thunder monster       |
| `dragon`        | Dragon monster        |
| `beast`         | Beast monster         |
| `beast_warrior` | Beast-Warrior monster |
| `dinosaur`      | Dinosaur monster      |
| `fish`          | Fish monster          |
| `sea_serpent`   | Sea Serpent monster   |
| `reptile`       | Reptile monster       |
| `psychic`       | Psychic monster       |
| `divine_beast`  | Divine-Beast monster  |
| `creator_god`   | Creator God monster   |
| `wyrm`          | Wyrm monster          |
| `cyberse`       | Cyberse monster       |


## Available attributes
| Attribute | Attribute for  |
| --------- | -------------- |
| `earth`   | Earth monster  |
| `water`   | Water monster  |
| `fire`    | Fire monster   |
| `wind`    | Wind monster   |
| `light`   | Light monster  |
| `dark`    | Dark monster   |
| `divine`  | Divine monster |


## Available categories
| Category            | Category for       |
| ------------------- | ------------------ |
| `destroy_monster`   | Destroy monster    |
| `destroy_spelltrap` | Destroy Spell/Trap |
| `destroy_deck`      | Destroy from deck  |
| `destroy_hand`      | Destroy from hand  |
| `send_to_gy`        | Sends to Graveyard |
| `send_to_hand`      | Sends to hand      |
| `send_to_deck`      | Sends to deck      |
| `banish`            | Banish             |
| `draw`              | Draws              |
| `search`            | Searches           |
| `change_atk_def`    | Changes ATK/DEF    |
| `change_level_rank` | Changes Level/Rank |
| `position`          | Changes position   |
| `piercing`          | Piercing damage    |
| `direct_attack`     | Direct attack      |
| `multi_attack`      | Multi attack       |
| `negate_activation` | Negates activation |
| `negate_effect`     | Negates effects    |
| `damage_lp`         | Damages LP         |
| `recover_lp`        | Recovers LP        |
| `special_summon`    | Special Summons    |
| `non_effect`        | Non-effect-related |
| `token_related`     | Token-related      |
| `fusion_related`    | Fusion-related     |
| `ritual_related`    | Ritual-related     |
| `synchro_related`   | Synchro-related    |
| `xyz_related`       | XYZ-related        |
| `link_related`      | Link-related       |
| `counter_related`   | Counter-related    |
| `gamble`            | Gamble             |
| `control`           | Control            |
| `move_zones`        | Move zones         |


## Available link arrows
| Link arrow     | Arrow direction |
| -------------- | --------------- |
| `bottom_left`  | Bottom-left     |
| `down`         | Down-center     |
| `bottom_right` | Bottom-right    |
| `left`         | Middle-left     |
| `right`        | Middle-right    |
| `upper_left`   | Top-left        |
| `top`          | Top-center      |
| `upper_right`  | Top-right       |
