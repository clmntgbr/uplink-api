<?php

declare(strict_types=1);

namespace App\Domain\Endpoint\Entity;

use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\GetCollection;
use ApiPlatform\Metadata\Post;
use App\Domain\Endpoint\Enum\MethodEnum;
use App\Domain\Endpoint\Repository\EndpointRepository;
use App\Domain\Project\Entity\Project;
use App\Domain\Step\Entity\Step;
use App\Infrastructure\Endpoint\Processor\CreateEndpointProcessor;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Symfony\Component\Serializer\Attribute\Groups;
use Symfony\Component\Uid\Uuid;

#[ORM\Entity(repositoryClass: EndpointRepository::class)]
#[ApiResource(
    order: ['createdAt' => 'DESC'],
    operations: [
        new GetCollection(
            normalizationContext: ['groups' => ['endpoint:read']],
        ),
        new Post(
            denormalizationContext: ['groups' => ['endpoint:write']],
            processor: CreateEndpointProcessor::class,
        ),
    ]
)]
class Endpoint
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING, nullable: true)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private ?string $name = null;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private string $baseUri;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private string $path;

    #[ORM\Column(type: Types::STRING, enumType: MethodEnum::class)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private MethodEnum $method = MethodEnum::GET;

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private array $header = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private array $body = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private array $response = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private array $query = [];

    #[ORM\Column(type: Types::BOOLEAN)]
    private bool $isActive = false;

    #[ORM\Column(type: Types::INTEGER, nullable: true)]
    #[Groups(['endpoint:read', 'endpoint:write', 'step:read'])]
    private ?int $timeoutSeconds = null;

    #[ORM\ManyToOne(targetEntity: Project::class, inversedBy: 'endpoints')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    private Project $project;

    /**
     * @var Collection<int, Step>
     */
    #[ORM\OneToMany(targetEntity: Step::class, mappedBy: 'endpoint', cascade: ['persist', 'remove'])]
    private Collection $steps;

    public function __construct()
    {
        $this->id = Uuid::v7();
        $this->steps = new ArrayCollection();
    }

    #[Groups(['endpoint:read'])]
    public function getId(): Uuid
    {
        return $this->id;
    }

    public function getMethod(): MethodEnum
    {
        return $this->method;
    }

    public function setMethod(MethodEnum $method): static
    {
        $this->method = $method;

        return $this;
    }

    /**
     * @return array<string, string>
     */
    public function getHeader(): array
    {
        return $this->header;
    }

    /**
     * @param array<string, string> $header
     */
    public function setHeader(array $header): static
    {
        $this->header = $header;

        return $this;
    }

    public function getBaseUri(): string
    {
        return $this->baseUri;
    }

    public function setBaseUri(string $baseUri): static
    {
        $this->baseUri = $baseUri;

        return $this;
    }

    public function getPath(): string
    {
        return $this->path;
    }

    public function setPath(string $path): static
    {
        $this->path = $path;

        return $this;
    }

    public function isActive(): bool
    {
        return $this->isActive;
    }

    public function setIsActive(bool $isActive): static
    {
        $this->isActive = $isActive;

        return $this;
    }

    /**
     * @return array<string, string>
     */
    public function getBody(): array
    {
        return $this->body;
    }

    /**
     * @param array<string, string> $body
     */
    public function setBody(array $body): static
    {
        $this->body = $body;

        return $this;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(?string $name): static
    {
        $this->name = $name;

        return $this;
    }

    public function getTimeoutSeconds(): ?int
    {
        return $this->timeoutSeconds;
    }

    public function setTimeoutSeconds(?int $timeoutSeconds): static
    {
        $this->timeoutSeconds = $timeoutSeconds;

        return $this;
    }

    public function getProject(): Project
    {
        return $this->project;
    }

    public function setProject(Project $project): static
    {
        $this->project = $project;

        return $this;
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
            $step->setEndpoint($this);
        }

        return $this;
    }

    public function removeStep(Step $step): static
    {
        $this->steps->removeElement($step);

        return $this;
    }

    /**
     * @return array<string, string>
     */
    public function getQuery(): array
    {
        return $this->query;
    }

    /**
     * @param array<string, string> $query
     */
    public function setQuery(array $query): static
    {
        $this->query = $query;

        return $this;
    }

    /**
     * @return array<string, string>
     */
    public function getResponse(): array
    {
        return $this->response;
    }

    /**
     * @param array<string, string> $response
     */
    public function setResponse(array $response): static
    {
        $this->response = $response;

        return $this;
    }
}
