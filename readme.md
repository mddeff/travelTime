# travelTime
A small Golang program that uses [Project OSRM](project-osrm.org) to calculate *driving* times between many sources and many destinations.

## **DISCLAIMER**
I wrote almost none of this.  This was largely copy/paste from a Large Language Model (LLM) output.  If you think this code violates a license you own/have the rights to, please don't hesitate to reach out.

## Why would I want this
This was born out of a need when shopping for property and I wanted to know how travel times from many sources (potential places to rent/purchase) were to many destinations (work, friends, family).  I couldn't find anything that would do it quite as easily, so here we are.

## Use

### Config
The config is relatively simple; a yaml file named `config.yaml` with two lists of dictionaries:

```yaml
sources:
- name: IAD
  location: "-77.45859148700659,38.95261987790932"
- name: "BWI"
  location: "-76.66849619071085,39.17783326145868"

destinations:
- name: LAS
  location: "-115.14829890428906,36.08328142666624"
- name: JFK
  location: "-73.77991713629551,40.65081823393867"
- name: AUS
  location: "-97.66671613141617,30.19546874439279"
```

**Note** - Order does matter; OSRM expects the input in `Longitude, Latitude`.  So if you're getting no distance returned, you might be trying to map a point on the other side of the world.

### Output
With the config above, you should see an output like this:
```
Source: IAD
    LAS: 42h58m, 2391.91Mi
    JFK: 5h33m, 262.82Mi
    AUS: 27h24m, 1521.48Mi
Source: BWI
    LAS: 43h9m, 2417.55Mi
    JFK: 4h26m, 211.45Mi
    AUS: 28h23m, 1570.55Mi
```

**Note** - Despite these being airports in the example above, this tool calculates *drive times*.

## Building/Running from source
You can get the latest release from the releases section, but you can also build it yourself:

```
go mod tidy
go build
```

Or you can run it directly from source:

```
go run main.go
```

