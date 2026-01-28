<?php

declare(strict_types=1);

namespace App\Domain\Step\Repository;

use App\Domain\Step\Entity\Step;
use App\Shared\Domain\Repository\AbstractRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends AbstractRepository<Step>
 */
class StepRepository extends AbstractRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Step::class);
    }
}
