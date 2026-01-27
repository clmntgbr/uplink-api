<?php

declare(strict_types=1);

namespace App\Domain\Endpoint\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Domain\Endpoint\Enum\MethodEnum;
use App\Domain\Endpoint\Repository\EndpointRepository;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;

#[ORM\Entity(repositoryClass: EndpointRepository::class)]
#[ApiResource]
class Endpoint
{
    use UuidTrait;
    use TimestampableEntity;

    #[ORM\Column(type: Types::STRING, nullable: true)]
    private ?string $name = null;

    #[ORM\Column(type: Types::STRING)]
    private string $baseUri;

    #[ORM\Column(type: Types::STRING)]
    private string $path;

    #[ORM\Column(type: Types::STRING, enumType: MethodEnum::class)]
    private MethodEnum $method = MethodEnum::GET;

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    private array $header = [];

    /**
     * @var array<string, string>
     */
    #[ORM\Column(type: Types::JSON)]
    private array $body = [];

    #[ORM\Column(type: Types::BOOLEAN)]
    private bool $isActive = true;

    #[ORM\Column(type: Types::INTEGER, nullable: true)]
    private ?int $timeoutSeconds = null;

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
}
