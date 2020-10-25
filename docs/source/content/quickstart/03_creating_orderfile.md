+++
title = "Creating new Orderfile.yml"
weight = 3
chapter = false
+++

{{% notice info %}}
Currently order doesn't support basic orderfile initialization
so new Orferfile.yml must be created by hand.\
However, this feature is planned in our [ROADMAP]({{< ref "/roadmap/#v020---gears-" >}} "ROADMAP") will be added very soon.
{{% /notice %}}


#### Creating new Orderfile.yml:
Create new file named `Orderfile.yml` containing structure like below
```YAML
version: "0.1.0"

orders:
  hello-world:
    description: "Say Hello World!"
    script:
      - echo "Hello World!"
```

This Orderfile contains single order named **"hello-world"**,
which can be invoked by running `order hello-world`.

