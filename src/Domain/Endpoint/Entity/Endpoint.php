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
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private ?string $name = null;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private string $baseUri;

    #[ORM\Column(type: Types::STRING)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private string $path;

    #[ORM\Column(type: Types::STRING, enumType: MethodEnum::class)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private MethodEnum $method = MethodEnum::GET;

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private array $header = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private array $body = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
    private array $param = [];

    #[ORM\Column(type: Types::BOOLEAN)]
    private bool $isActive = false;

    #[ORM\Column(type: Types::INTEGER, nullable: true)]
    #[Groups(['endpoint:read', 'endpoint:write'])]
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

    public function setMethod(MethodEnum $method): void
    {
        $this->method = $method;
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
    public function setHeader(array $header): void
    {
        $this->header = $header;
    }

    public function getBaseUri(): string
    {
        return $this->baseUri;
    }

    public function setBaseUri(string $baseUri): void
    {
        $this->baseUri = $baseUri;
    }

    public function getPath(): string
    {
        return $this->path;
    }

    public function setPath(string $path): void
    {
        $this->path = $path;
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
     * @return array<string, string>
     */
    public function getBody(): array
    {
        return $this->body;
    }

    /**
     * @param array<string, string> $body
     */
    public function setBody(array $body): void
    {
        $this->body = $body;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(?string $name): void
    {
        $this->name = $name;
    }

    public function getTimeoutSeconds(): ?int
    {
        return $this->timeoutSeconds;
    }

    public function setTimeoutSeconds(?int $timeoutSeconds): void
    {
        $this->timeoutSeconds = $timeoutSeconds;
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

    public function addStep(Step $step): void
    {
        if (! $this->steps->contains($step)) {
            $this->steps->add($step);
            $step->setEndpoint($this);
        }
    }

    public function removeStep(Step $step): void
    {
        $this->steps->removeElement($step);
    }

    public function getParam(): array
    {
        return $this->param;
    }

    public function setParam(array $param): static
    {
        $this->param = $param;

        return $this;
    }
}
