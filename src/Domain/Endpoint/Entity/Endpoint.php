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

    #[ORM\Column(type: Types::STRING)]
    private string $baseUri;

    #[ORM\Column(type: Types::STRING)]
    private string $path;

    #[ORM\Column(type: Types::STRING, enumType: MethodEnum::class)]
    private MethodEnum $method = MethodEnum::GET;

    /**
     * @var array<string, string>|null
     */
    #[ORM\Column(type: Types::JSON)]
    private ?array $headers = [];

    #[ORM\Column(type: Types::BOOLEAN)]
    private bool $isActive = true;

    public function getMethod(): MethodEnum
    {
        return $this->method;
    }

    public function setMethod(MethodEnum $method): void
    {
        $this->method = $method;
    }

    /**
     * @return array<string, string>|null
     */
    public function getHeaders(): ?array
    {
        return $this->headers;
    }

    /**
     * @param array<string, string>|null $headers
     */
    public function setHeaders(?array $headers): void
    {
        $this->headers = $headers;
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
}
