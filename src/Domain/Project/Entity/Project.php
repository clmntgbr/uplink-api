<?php

declare(strict_types=1);

namespace App\Domain\Project\Entity;

use Doctrine\Common\Collections\ArrayCollection;
use ApiPlatform\Metadata\ApiResource;
use App\Domain\Endpoint\Entity\Endpoint;
use App\Domain\Project\Repository\ProjectRepository;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;

#[ORM\Entity(repositoryClass: ProjectRepository::class)]
#[ApiResource]
class Project
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING)]
    private string $name;

    #[ORM\Column(type: Types::BOOLEAN)]
    private bool $isActive = false;

    /**
     * @var Collection<int, Endpoint>
     */
    #[ORM\OneToMany(targetEntity: Endpoint::class, mappedBy: 'project', cascade: ['persist', 'remove'])]
    private Collection $endpoints;

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
    public function __construct()
    {
        $this->endpoints = new ArrayCollection();
    }
}
