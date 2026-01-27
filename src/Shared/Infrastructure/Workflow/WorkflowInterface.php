<?php

declare(strict_types=1);

namespace App\Shared\Infrastructure\Workflow;

use App\Domain\Clip\Entity\Clip;

interface WorkflowInterface
{
    public function canApply(Clip $clip, string $transition): bool;

    public function apply(Clip $clip, string $transition): void;
}
