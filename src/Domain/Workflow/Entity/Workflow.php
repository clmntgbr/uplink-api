<?php

declare(strict_types=1);

namespace App\Domain\Workflow\Entity;

use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\Get;
use ApiPlatform\Metadata\GetCollection;
use ApiPlatform\Metadata\Post;
use App\Domain\Project\Entity\Project;
use App\Domain\Step\Entity\Step;
use App\Domain\Workflow\Repository\WorkflowRepository;
use App\Infrastructure\Project\Validation\Constraint\MaxWorkflowsPerProject;
use App\Infrastructure\Workflow\Processor\CreateWorkflowProcessor;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Symfony\Component\Serializer\Attribute\Groups;
use Symfony\Component\Uid\Uuid;
use Symfony\Component\Validator\Constraints as Assert;

#[ORM\Entity(repositoryClass: WorkflowRepository::class)]
#[ApiResource(
    operations: [
        new GetCollection(
            normalizationContext: ['groups' => ['workflow:read', 'step:read']],
        ),
        new Post(
            denormalizationContext: ['groups' => ['workflow:write']],
            validationContext: ['groups' => [MaxWorkflowsPerProject::GROUP_CREATE]],
            processor: CreateWorkflowProcessor::class,
        ),
        new Get(
            normalizationContext: ['groups' => ['workflow:read', 'step:read']],
        ),
    ]
)]
class Workflow
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['workflow:read', 'workflow:write'])]
    #[Assert\NotBlank]
    #[Assert\Length(min: 3, max: 255)]
    private string $name;

    #[ORM\Column(type: Types::TEXT, nullable: true)]
    #[Groups(['workflow:read', 'workflow:write'])]
    private ?string $description = null;

    /**
     * @var Collection<int, Step>
     */
    #[ORM\OneToMany(targetEntity: Step::class, mappedBy: 'workflow', cascade: ['persist', 'remove'])]
    private Collection $steps;

    #[ORM\ManyToOne(targetEntity: Project::class, inversedBy: 'workflows')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    #[MaxWorkflowsPerProject(groups: [MaxWorkflowsPerProject::GROUP_CREATE])]
    private Project $project;

    public function __construct()
    {
        $this->id = Uuid::v7();
        $this->steps = new ArrayCollection();
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

    public function getDescription(): ?string
    {
        return $this->description;
    }

    public function setDescription(?string $description): void
    {
        $this->description = $description;
    }

    public function getProject(): Project
    {
        return $this->project;
    }

    public function setProject(Project $project): void
    {
        $this->project = $project;
    }

    /**
     * @return Collection<int, Step>
     */
    public function getSteps(): Collection
    {
        return $this->steps;
    }

    public function addStep(Step $step): static
    {
        if (! $this->steps->contains($step)) {
            $this->steps->add($step);
            $step->setWorkflow($this);
        }

        return $this;
    }

    public function removeStep(Step $step): static
    {
        $this->steps->removeElement($step);

        return $this;
    }
}
