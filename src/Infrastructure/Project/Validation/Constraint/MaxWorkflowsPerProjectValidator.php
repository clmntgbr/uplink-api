<?php

declare(strict_types=1);

namespace App\Infrastructure\Project\Validation\Constraint;

use App\Domain\Project\Entity\Project;
use Override;
use Symfony\Component\Validator\Constraint;
use Symfony\Component\Validator\ConstraintValidator;
use Symfony\Component\Validator\Exception\UnexpectedTypeException;
use Symfony\Component\Validator\Exception\UnexpectedValueException;

final class MaxWorkflowsPerProjectValidator extends ConstraintValidator
{
    #[Override]
    public function validate(mixed $value, Constraint $constraint): void
    {
        if (! $constraint instanceof MaxWorkflowsPerProject) {
            throw new UnexpectedTypeException($constraint, MaxWorkflowsPerProject::class);
        }

        if (null === $value) {
            return;
        }

        if (! $value instanceof Project) {
            throw new UnexpectedValueException($value, Project::class);
        }

        if ($value->getWorkflows()->count() >= MaxWorkflowsPerProject::MAX_WORKFLOWS) {
            $this->context->buildViolation($constraint->message)
                ->setCode(MaxWorkflowsPerProject::CODE)
                ->setParameter('{{ limit }}', (string) MaxWorkflowsPerProject::MAX_WORKFLOWS)
                ->addViolation();
        }
    }
}
