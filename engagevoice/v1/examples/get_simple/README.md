# Engage Voice Go Client - Getting Started

## Call the `countries` API to get a list of available Countries

```
$ main.go -e .env -o country
```

## Call the `agentgroup` API to get a list of Agent Groups

```
$ main.go -e .env -o agentgroup
```

This will return a list of agentGroupIds.

## Call the `agent` API to get a list of Agents for an Agent Group

```
$ main.go -e .env -o agent -i <agentGroupId>
```

This will return a list of agents.
