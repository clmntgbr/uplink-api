<?php

declare(strict_types=1);

namespace App\Application\User\Command;

use App\Domain\User\Entity\User;
use App\Domain\User\Repository\UserRepository;
use App\Shared\Application\Bus\CommandBusInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class CreateOrUpdateUserCommandHandler
{
    public function __construct(
        private readonly UserRepository $userRepository,
        private readonly CommandBusInterface $commandBus,
    ) {
    }

    public function __invoke(CreateOrUpdateUserCommand $command): User
    {
        $user = $this->userRepository->findOneBy(['email' => $command->getEmail()]);

        if (null === $user) {
            $user = $this->commandBus->dispatch(new CreateUserCommand(
                email: $command->getEmail(),
                plainPassword: $command->getPlainPassword(),
                firstname: $command->getFirstname(),
                lastname: $command->getLastname(),
                picture: $command->getPicture(),
            ));
        }

        return $user;
    }
}
