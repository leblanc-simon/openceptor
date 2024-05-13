<?php

namespace App\Entity;

use App\Repository\ProjectRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Bridge\Doctrine\IdGenerator\UuidGenerator;
use Symfony\Bridge\Doctrine\Validator\Constraints\UniqueEntity;

#[ORM\Entity(repositoryClass: ProjectRepository::class)]
#[UniqueEntity(fields: ['endpoint'])]
class Project
{
    #[ORM\Id]
    #[ORM\Column(type: 'uuid', unique:true)]
    #[ORM\GeneratedValue(strategy: 'CUSTOM')]
    #[ORM\CustomIdGenerator(class: UuidGenerator::class)]
    private string $id;

    #[ORM\Column(type: 'string', length: 100)]
    private string $name;

    #[ORM\Column(type: 'string', length: 36, unique: true)]
    private string $endpoint;

    #[ORM\OneToMany(mappedBy: 'project', targetEntity: MockingRule::class, orphanRemoval: true)]
    private Collection $mockingRules;

    public function __construct()
    {
        $this->mockingRules = new ArrayCollection();
    }

    public function getId(): ?string
    {
        return $this->id;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(string $name): self
    {
        $this->name = $name;

        return $this;
    }

    public function getEndpoint(): ?string
    {
        return $this->endpoint;
    }

    public function setEndpoint(string $endpoint): self
    {
        $this->endpoint = $endpoint;

        return $this;
    }

    /**
     * @return Collection<int, MockingRule>
     */
    public function getMockingRules(): Collection
    {
        return $this->mockingRules;
    }

    public function addMockingRule(MockingRule $mockingRule): self
    {
        if (!$this->mockingRules->contains($mockingRule)) {
            $this->mockingRules->add($mockingRule);
            $mockingRule->setProject($this);
        }

        return $this;
    }

    public function removeMockingRule(MockingRule $mockingRule): self
    {
        if ($this->mockingRules->removeElement($mockingRule)) {
            // set the owning side to null (unless already changed)
            if ($mockingRule->getProject() === $this) {
                $mockingRule->setProject(null);
            }
        }

        return $this;
    }
}
