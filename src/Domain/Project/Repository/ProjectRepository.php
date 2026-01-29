<?php

declare(strict_types=1);

namespace App\Domain\Project\Repository;

use App\Domain\Project\Entity\Project;
use App\Shared\Domain\Repository\AbstractRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends AbstractRepository<Project>
 */
class ProjectRepository extends AbstractRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Project::class);
    }
}
