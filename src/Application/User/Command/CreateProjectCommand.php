<?php

declare(strict_types=1);

namespace App\Application\User\Command;

use App\Domain\User\Entity\User;
use App\Shared\Application\Command\SynchronousInterface;

final class CreateProjectCommand implements SynchronousInterface
{
    public function __construct(
        public string $name,
        public User $user,
        public bool $isActive = false,
    ) {
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function getUser(): User
    {
        return $this->user;
    }

    public function isActive(): bool
    {
        return $this->isActive;
    }
}
