<?php

declare(strict_types=1);

namespace App\Domain\Workflow\Repository;

use App\Domain\Workflow\Entity\Workflow;
use App\Shared\Domain\Repository\AbstractRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends AbstractRepository<Workflow>
 */
class WorkflowRepository extends AbstractRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Workflow::class);
    }
}
