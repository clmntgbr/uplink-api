<?php

declare(strict_types=1);

namespace App\Domain\Workflow\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Domain\Workflow\Repository\WorkflowRepository;
use App\Shared\Domain\Trait\UuidTrait;
use Doctrine\ORM\Mapping as ORM;
use Gedmo\Timestampable\Traits\TimestampableEntity;

#[ORM\Entity(repositoryClass: WorkflowRepository::class)]
#[ApiResource]
class Workflow
{
    use UuidTrait;
    use TimestampableEntity;
}
