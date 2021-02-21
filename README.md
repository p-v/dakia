dakia
=====

Remember commands and execute them later

#### Store a command

A command can be saved in as a key value pair

```
# dakia -s <alias> -- <command to save>
dakia -s list parent directory -- ls -l ../
```

Commands can be nested

```
dakia -s list this directory -- ls -l
```

#### Find saved commands

To find all top level commands

```
# dakia -f [alias]
dakia -f
```

To find nested commands

```
dakia -f list
```

#### Executing a command

Any store command can be executed as,

```
# dakia <saved alias>
dakia list this directory
```
