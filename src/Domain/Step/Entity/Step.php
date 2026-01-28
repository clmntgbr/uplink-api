<?php

declare(strict_types=1);

namespace App\Domain\Step\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Domain\Endpoint\Entity\Endpoint;
use App\Domain\Step\Repository\StepRepository;
use App\Domain\Workflow\Entity\Workflow;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;

#[ORM\Entity(repositoryClass: StepRepository::class)]
#[ApiResource]
class Step
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::INTEGER)]
    private int $position;

    #[ORM\ManyToOne(targetEntity: Endpoint::class, inversedBy: 'steps')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    private Endpoint $endpoint;

    #[ORM\ManyToOne(targetEntity: Workflow::class, inversedBy: 'steps')]
    #[ORM\JoinColumn(nullable: false, onDelete: 'CASCADE')]
    private Workflow $workflow;

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    private array $variables = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    private array $outputs = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    private array $asserts = [];

    /**
     * @return array<string, string>
     */
    public function getVariables(): array
    {
        return $this->variables;
    }

    /**
     * @param array<string, string> $variables
     */
    public function setVariables(array $variables): void
    {
        $this->variables = $variables;
    }

    /**
     * @return array<string, string>
     */
    public function getOutputs(): array
    {
        return $this->outputs;
    }

    /**
     * @param array<string, string> $outputs
     */
    public function setOutputs(array $outputs): void
    {
        $this->outputs = $outputs;
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
    public function getAsserts(): array
    {
        return $this->asserts;
    }

    /**
     * @param array<string, string> $asserts
     */
    public function setAsserts(array $asserts): void
    {
        $this->asserts = $asserts;
    }
}
