<?php

declare(strict_types=1);

namespace App\Domain\Project\Entity;

use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\GetCollection;
use ApiPlatform\Metadata\Patch;
use ApiPlatform\Metadata\Post;
use App\Domain\Endpoint\Entity\Endpoint;
use App\Domain\Project\Repository\ProjectRepository;
use App\Domain\User\Entity\User;
use App\Domain\Workflow\Entity\Workflow;
use App\Infrastructure\Project\Processor\CreateProjectProcessor;
use App\Infrastructure\Project\Processor\UpdateProjectProcessor;
use App\Infrastructure\Project\Validation\Constraint\MaxProjectsPerUser;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Symfony\Component\Serializer\Attribute\Groups;
use Symfony\Component\Uid\Uuid;

#[ORM\Entity(repositoryClass: ProjectRepository::class)]
#[ApiResource(
    operations: [
        new GetCollection(
            normalizationContext: ['groups' => ['project:read']],
        ),
        new Post(
            denormalizationContext: ['groups' => ['project:write']],
            validationContext: ['groups' => ['Default', MaxProjectsPerUser::GROUP_CREATE]],
            processor: CreateProjectProcessor::class,
        ),
        new Patch(
            denormalizationContext: ['groups' => ['project:write']],
            processor: UpdateProjectProcessor::class,
        ),
    ]
)]
class Project
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['project:read', 'project:write'])]
    private string $name;

    #[ORM\ManyToOne(targetEntity: User::class, inversedBy: 'projects')]
    #[ORM\JoinColumn(nullable: false)]
    #[MaxProjectsPerUser(groups: [MaxProjectsPerUser::GROUP_CREATE])]
    private User $user;

    /**
     * @var Collection<int, Endpoint>
     */
    #[ORM\OneToMany(targetEntity: Endpoint::class, mappedBy: 'project', cascade: ['persist', 'remove'])]
    private Collection $endpoints;

    /**
     * @var Collection<int, Workflow>
     */
    #[ORM\OneToMany(targetEntity: Workflow::class, mappedBy: 'project', cascade: ['persist', 'remove'])]
    private Collection $workflows;

    public function __construct()
    {
        $this->id = Uuid::v7();
        $this->endpoints = new ArrayCollection();
        $this->workflows = new ArrayCollection();
    }

    #[Groups(['project:read'])]
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

    #[Groups(['project:write'])]
    public function setActive(bool $active): self
    {
        if ($active) {
            $this->user->setActiveProject($this);
        }

        return $this;
    }

    /**
     * @return Collection<int, Endpoint>
     */
    public function getEndpoints(): Collection
    {
        return $this->endpoints;
    }

    /**
     * @param Collection<int, Endpoint> $endpoints
     */
    public function setEndpoints(Collection $endpoints): void
    {
        $this->endpoints = $endpoints;
    }

    public function addEndpoint(Endpoint $endpoint): void
    {
        $this->endpoints->add($endpoint);
    }

    public function removeEndpoint(Endpoint $endpoint): void
    {
        $this->endpoints->removeElement($endpoint);
    }

    public function getUser(): User
    {
        return $this->user;
    }

    public function setUser(User $user): void
    {
        $this->user = $user;
    }

    /**
     * @return Collection<int, Workflow>
     */
    public function getWorkflows(): Collection
    {
        return $this->workflows;
    }

    public function addWorkflow(Workflow $workflow): void
    {
        if (! $this->workflows->contains($workflow)) {
            $this->workflows->add($workflow);
            $workflow->setProject($this);
        }
    }

    public function removeWorkflow(Workflow $workflow): void
    {
        $this->workflows->removeElement($workflow);
    }

    #[Groups(['project:read'])]
    public function getIsActive(): bool
    {
        return $this->user->getActiveProject() === $this;
    }
}
