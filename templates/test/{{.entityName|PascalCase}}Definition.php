<?php declare(strict_types=1);

namespace {{.namespace}};

use Shopware\Core\Framework\DataAbstractionLayer\Field\Flag\Required;
use Shopware\Core\Framework\DataAbstractionLayer\EntityDefinition;
use Shopware\Core\Framework\DataAbstractionLayer\Field\Flag\PrimaryKey;
use Shopware\Core\Framework\DataAbstractionLayer\Field\IdField;
use Shopware\Core\Framework\DataAbstractionLayer\FieldCollection;
use Shopware\Core\Framework\DataAbstractionLayer\Field\Flag\ApiAware;
use Shopware\Core\Framework\DataAbstractionLayer\Field\TranslationsAssociationField;
use {{.namespace}}\Aggregate\Translation\{{.entityName|PascalCase}}TranslationDefinition;


class {{.entityName|PascalCase}}Definition extends EntityDefinition
{
    public const ENTITY_NAME = '{{.tableName|SnakeCase}}';

    public function getEntityName(): string
    {
        return self::ENTITY_NAME;
    }

    public function getEntityClass(): string
    {
        return {{.entityName|PascalCase}}Entity::class;
    }

    public function getCollectionClass(): string
    {
        return {{.entityName|PascalCase}}Collection::class;
    }

    protected function defineFields(): FieldCollection
    {
        return new FieldCollection([
            (new IdField('id', 'id'))->addFlags(new Required(), new PrimaryKey()),
            (new TranslationsAssociationField(
                {{.entityName|PascalCase}}TranslationDefinition::class,
                '{{.entityName|SnakeCase}}_id'
            ))->addFlags(new ApiAware(), new Required())
        ]);
    }
}