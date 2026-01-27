<?php

declare(strict_types=1);

namespace App\Shared\Infrastructure\Workflow;

use App\Domain\Clip\Entity\Clip;
use App\Shared\Infrastructure\Workflow\WorkflowInterface as SharedWorkflowInterface;
use Override;
use RuntimeException;
use Symfony\Component\Workflow\WorkflowInterface;

use function sprintf;

class Workflow implements SharedWorkflowInterface
{
    public function __construct(
        private readonly WorkflowInterface $clipsStateMachine,
    ) {
    }

    #[Override]
    public function canApply(Clip $clip, string $transition): bool
    {
        return $this->clipsStateMachine->can($clip, $transition);
    }

    #[Override]
    public function apply(Clip $clip, string $transition): void
    {
        if (false === $this->canApply($clip, $transition)) {
            throw new RuntimeException(sprintf('Transition "%s" cannot be applied to clip "%s"', $transition, $clip->getId()));
        }

        $this->clipsStateMachine->apply($clip, $transition);
    }
}
