<?php

declare(strict_types=1);

namespace App\Domain\Project\Entity;

use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\GetCollection;
use App\Domain\Endpoint\Entity\Endpoint;
use App\Domain\Project\Repository\ProjectRepository;
use App\Domain\User\Entity\User;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Symfony\Component\Uid\Uuid;

#[ORM\Entity(repositoryClass: ProjectRepository::class)]
#[ApiResource(
    operations: [
        new GetCollection(),
    ]
)]
class Project
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING)]
    private string $name;

    #[ORM\Column(type: Types::BOOLEAN)]
    private bool $isActive = false;

    #[ORM\ManyToOne(targetEntity: User::class, inversedBy: 'projects')]
    #[ORM\JoinColumn(nullable: false)]
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

    public function setIsActive(bool $isActive): void
    {
        $this->isActive = $isActive;
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
}
