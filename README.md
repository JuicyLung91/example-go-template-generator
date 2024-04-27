# Introduction
This is just a quick example of how to generate boilerplate code with generic template files in go.


## Templates

The templates are working with go text/template https://pkg.go.dev/text/template
The `text/template` package provides a powerful templating engine that is easy to use and flexible.

The example templates are located in `templates/test`.
The templates are working recursively with subdirectories. You can add template for directories and for files.

### Example template
```
<?php declare(strict_types=1);

use Shopware\Core\Framework\DataAbstractionLayer\EntityDefinition;


namespace {{.namespace}};

class {{.entityName|PascalCase}}Definition extends EntityDefinition
{ 
    public const ENTITY_NAME = '{{.tableName|SnakeCase}}';
}
```

Will result in
```
<?php declare(strict_types=1);

namespace Swag\MyFancyPlugin\Content\Entity\SwagMyFancyEntity;

use Shopware\Core\Framework\DataAbstractionLayer\EntityDefinition;
class SwagMyFancyEntityDefinition extends EntityDefinition
{
    public const ENTITY_NAME = 'swag_my_fancy_entity';
}
```

### Variables
Variables can be defined and accessed within templates. Here's how you can define and use variables:

```
{{ .GlobalVariableName | SnakeCase }}

//or static variables

{{ $variable := "Hello, World!" }}
{{ $variable }}
```

### 

### If Queries
You can use if queries to conditionally render content based on boolean conditions. Here's an example:

```
{{ if .Condition }}
    This content will be rendered if Condition is true.
{{ else }}
    This content will be rendered if Condition is false.
{{ end }}
```


### Loops
You can use loops to iterate over collections such as slices, arrays, maps, or even strings. Here's an example of looping over a slice:

```
{{ range .Items }}
    {{ . }} <!-- Access each item in the slice -->
{{ end }}
```


### Functions
Go templates support function calls for more complex operations. You can define custom functions or use built-in functions. Here's an example of using a built-in function:

```
{{ $result := add 2 3 }}
{{ $result }} <!-- Output: 5 -->
```


## Case Transformation Methods

The following helper methods have been provided flexibility in transforming string cases according to your specific requirements.

### Camel Case
Camel case is a naming convention where the first letter of each word except the first is capitalized and spaces are removed. Here's how you can convert a string to camel case:

`{{ .Variable | CamelCase }}`

### Snake Case
Snake case is a naming convention where words are separated by underscores and all letters are lowercase. Here's how you can convert a string to snake case:

`{{ .Variable | SnakeCase }}`

### Kebab Case
Kebab case is similar to snake case, but words are separated by hyphens instead of underscores. Here's how you can convert a string to kebab case:

`{{ .Variable | KebabCase }}`

### Pascal Case
Pascal case is similar to camel case, but the first letter of each word is capitalized. Here's how you can convert a string to Pascal case:

`{{ .Variable | PascalCase }}`




## Run a test
The docker compose file just runs the `hello.go` script. 
Just run `docker-compose up` and check the  `testoutput` directory.
You can adjust variables for the template inside the `hello.go` main function. You can add more file inside the template directory if you want. 