<?php

declare(strict_types=1);

namespace App\Domain\Step\Dto;

class UpdateStepItem
{
    public function __construct(
        public string $id,
        public int $position,
    ) {
    }

    public function getId(): string
    {
        return $this->id;
    }

    public function getPosition(): int
    {
        return $this->position;
    }
}
