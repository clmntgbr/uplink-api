<?php

declare(strict_types=1);

namespace App\Domain\Endpoint\Repository;

use App\Domain\Endpoint\Entity\Endpoint;
use App\Shared\Domain\Repository\AbstractRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends AbstractRepository<Endpoint>
 */
class EndpointRepository extends AbstractRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Endpoint::class);
    }
}
