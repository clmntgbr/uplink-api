<?php

declare(strict_types=1);

namespace App\Application\User\Command;

use App\Domain\User\Entity\User;
use App\Domain\User\Repository\UserRepository;
use App\Shared\Application\Bus\CommandBusInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class CreateUserCommandHandler
{
    public function __construct(
        private readonly UserRepository $userRepository,
        private readonly CommandBusInterface $commandBus,

    ) {
    }

    public function __invoke(CreateUserCommand $command): User
    {
        $user = new User();
        $user->setPlainPassword($command->getPlainPassword());
        $user->setFirstname($command->getFirstname());
        $user->setLastname($command->getLastname());
        $user->setPicture($command->getPicture());
        $user->setEmail($command->getEmail());

        $this->userRepository->save($user, true);

        $this->commandBus->dispatch(new CreateProjectCommand(
            name: 'Default project',
            user: $user,
            isActive: true,
        ));

        return $user;
    }
}
