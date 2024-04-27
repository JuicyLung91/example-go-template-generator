<?php declare(strict_types=1);

namespace {{.namespace}};

use Shopware\Core\Framework\DataAbstractionLayer\EntityTranslationDefinition;
use Shopware\Core\Framework\DataAbstractionLayer\FieldCollection;
use {{.parentClassNamespace}}Definition;

class {{.entityName|PascalCase}}TranslationDefinition extends EntityTranslationDefinition
{
    public const ENTITY_NAME = '{{.tableName|SnakeCase}}_translation';

    public function getEntityName(): string
    {
        return self::ENTITY_NAME;
    }

    public function getParentDefinitionClass(): string
    {
        return {{.entityName|PascalCase}}Definition::class;
    }

    public function getEntityClass(): string
    {
        return {{.entityName|PascalCase}}TranslationEntity::class;
    }

    public function getCollectionClass(): string
    {
        return {{.entityName|PascalCase}}TranslationCollection::class;
    }

    protected function defineFields(): FieldCollection
    {
        return new FieldCollection([

        ]);
    }
}