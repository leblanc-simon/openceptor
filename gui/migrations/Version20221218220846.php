<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20221218220846 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE TABLE mocking_rule (id UUID NOT NULL, project_id UUID NOT NULL, method VARCHAR(7) NOT NULL, condition_type VARCHAR(255) NOT NULL, path VARCHAR(255) NOT NULL, response_status INT NOT NULL, response_headers JSON NOT NULL, response_body TEXT DEFAULT NULL, description VARCHAR(255) DEFAULT NULL, created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL, updated_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL, PRIMARY KEY(id))');
        $this->addSql('CREATE INDEX IDX_36E1F17D166D1F9C ON mocking_rule (project_id)');
        $this->addSql('COMMENT ON COLUMN mocking_rule.id IS \'(DC2Type:uuid)\'');
        $this->addSql('COMMENT ON COLUMN mocking_rule.project_id IS \'(DC2Type:uuid)\'');
        $this->addSql('COMMENT ON COLUMN mocking_rule.created_at IS \'(DC2Type:datetime_immutable)\'');
        $this->addSql('COMMENT ON COLUMN mocking_rule.updated_at IS \'(DC2Type:datetime_immutable)\'');
        $this->addSql('ALTER TABLE mocking_rule ADD CONSTRAINT FK_36E1F17D166D1F9C FOREIGN KEY (project_id) REFERENCES project (id) NOT DEFERRABLE INITIALLY IMMEDIATE');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE SCHEMA public');
        $this->addSql('ALTER TABLE mocking_rule DROP CONSTRAINT FK_36E1F17D166D1F9C');
        $this->addSql('DROP TABLE mocking_rule');
    }
}
