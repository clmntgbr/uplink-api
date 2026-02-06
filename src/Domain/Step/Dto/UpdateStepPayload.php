<?php

declare(strict_types=1);

namespace App\Domain\Step\Dto;

class UpdateStepPayload
{
    /**
     * @param UpdateStepItem[] $steps
     */
    public function __construct(
        public string $workflow,
        public array $steps,
    ) {
    }
}
