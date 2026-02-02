<?php

declare(strict_types=1);

namespace App\Application\User\Command;

use App\Domain\Project\Entity\Project;
use App\Domain\Project\Repository\ProjectRepository;
use App\Domain\User\Entity\User;
use App\Domain\User\Repository\UserRepository;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class CreateProjectCommandHandler
{
    public function __construct(
        private readonly ProjectRepository $projectRepository,
    ) {
    }

    public function __invoke(CreateProjectCommand $command): Project
    {
        $project = new Project();
        $project->setName($command->getName());
        $project->setUser($command->getUser());
        $project->setIsActive($command->isActive());

        $this->projectRepository->save($project, true);

        return $project;
    }
}
