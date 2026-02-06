<?php

declare(strict_types=1);

namespace App\Domain\Step\Entity;

use ApiPlatform\Doctrine\Orm\Filter\SearchFilter;
use ApiPlatform\Metadata\ApiFilter;
use ApiPlatform\Metadata\ApiResource;
use ApiPlatform\Metadata\GetCollection;
use ApiPlatform\Metadata\Patch;
use ApiPlatform\Metadata\Post;
use ApiPlatform\Metadata\QueryParameter;
use App\Domain\Endpoint\Entity\Endpoint;
use App\Domain\Step\Repository\StepRepository;
use App\Domain\Workflow\Entity\Workflow;
use App\Infrastructure\Step\Processor\CreateStepProcessor;
use App\Infrastructure\Project\Validation\Constraint\MaxStepsPerWorkflow;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;
use Symfony\Component\Serializer\Attribute\Groups;
use Symfony\Component\Uid\Uuid;

#[ORM\Entity(repositoryClass: StepRepository::class)]
#[ApiResource(
    order: ['position' => 'ASC'],
    operations: [
        new GetCollection(
            normalizationContext: ['groups' => ['step:read']],
            strictQueryParameterValidation: true,
            parameters: [
                'workflow' => new QueryParameter(description: 'Filter steps by workflow'),
                'itemsPerPage' => new QueryParameter(description: 'Items per page'),
            ]
        ),
        new Post(
            denormalizationContext: ['groups' => ['step:write']],
            validationContext: ['groups' => [MaxStepsPerWorkflow::GROUP_CREATE]],
            processor: CreateStepProcessor::class,
        ),
        new Patch(
            denormalizationContext: ['groups' => ['step:write']],
        ),
    ]
)]
#[ApiFilter(SearchFilter::class, properties: ['workflow' => 'exact'])]
class Step
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::INTEGER)]
    #[Groups(['step:read'])]
    private int $position;

    #[ORM\ManyToOne(targetEntity: Endpoint::class, inversedBy: 'steps')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    #[Groups(['step:read', 'step:write'])]
    private Endpoint $endpoint;

    #[ORM\ManyToOne(targetEntity: Workflow::class, inversedBy: 'steps')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    #[MaxStepsPerWorkflow(groups: [MaxStepsPerWorkflow::GROUP_CREATE])]
    #[Groups(['step:write'])]
    private Workflow $workflow;

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'step:write', 'step:read'])]
    private array $header = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'step:write', 'step:read'])]
    private array $body = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'step:write', 'step:read'])]
    private array $response = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    #[Groups(['endpoint:read', 'step:write', 'step:read'])]
    private array $query = [];

    public function __construct()
    {
        $this->id = Uuid::v7();
    }

    #[Groups(['step:read'])]
    public function getId(): Uuid
    {
        return $this->id;
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
    public function setResponse(array $response): void
    {
        $this->response = $response;
    }

    public function getPosition(): int
    {
        return $this->position;
    }

    public function setPosition(int $position): void
    {
        $this->position = $position;
    }

    public function getEndpoint(): Endpoint
    {
        return $this->endpoint;
    }

    public function setEndpoint(Endpoint $endpoint): void
    {
        $this->endpoint = $endpoint;
    }

    public function getWorkflow(): Workflow
    {
        return $this->workflow;
    }

    public function setWorkflow(Workflow $workflow): void
    {
        $this->workflow = $workflow;
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
}
