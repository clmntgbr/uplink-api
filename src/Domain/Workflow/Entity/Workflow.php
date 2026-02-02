<?php

declare(strict_types=1);

namespace App\Domain\Workflow\Entity;

use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\GetCollection;
use App\Domain\Project\Entity\Project;
use App\Domain\Workflow\Repository\WorkflowRepository;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Symfony\Component\Serializer\Attribute\Groups;
use Symfony\Component\Uid\Uuid;

#[ORM\Entity(repositoryClass: WorkflowRepository::class)]
#[ApiResource(
    operations: [
        new GetCollection(
            normalizationContext: ['groups' => ['workflow:read']],
        ),
    ]
)]
class Workflow
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['workflow:read'])]
    private string $name;

    #[ORM\ManyToOne(targetEntity: Project::class, inversedBy: 'workflows')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    private Project $project;

    public function __construct()
    {
        $this->id = Uuid::v7();
    }

    #[Groups(['workflow:read'])]
    public function getId(): Uuid
    {
        return $this->id;
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function setName(string $name): void
    {
        $this->name = $name;
    }

    public function getProject(): Project
    {
        return $this->project;
    }

    public function setProject(Project $project): void
    {
        $this->project = $project;
    }
}
