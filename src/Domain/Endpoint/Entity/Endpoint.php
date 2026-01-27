<?php

declare(strict_types=1);

namespace App\Domain\Endpoint\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Domain\Endpoint\Repository\EndpointRepository;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;

#[ORM\Entity(repositoryClass: EndpointRepository::class)]
#[ApiResource]
class Endpoint
{
    use UuidTrait;
    use TimestampableEntity;
}
