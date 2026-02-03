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

    #[ORM\Column(type: Types::BOOLEAN)]
    #[Groups(['project:read'])]
    private bool $isActive = false;

    #[ORM\ManyToOne(targetEntity: User::class, inversedBy: 'projects')]
    #[ORM\JoinColumn(nullable: false)]
    #[MaxProjectsPerUser]
    private User $user;

    /**
     * @var Collection<int, Endpoint>
     */
    #[ORM\OneToMany(targetEntity: Endpoint::class, mappedBy: 'project', cascade: ['persist', 'remove'])]
    private Collection $endpoints;

    public function __construct()
    {
        $this->id = Uuid::v7();
        $this->endpoints = new ArrayCollection();
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

    public function isActive(): bool
    {
        return $this->isActive;
    }

    public function setIsActive(bool $isActive): self
    {
        $this->isActive = $isActive;

        return $this;
    }

    #[Groups(['project:write'])]
    public function setActive(bool $isActive): self
    {
        $this->isActive = $isActive;

        if ($isActive) {
            $this->user->updateProjects($this);
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

    public function getIsActive(): bool
    {
        return $this->isActive;
    }
}
