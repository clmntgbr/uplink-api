<?php

declare(strict_types=1);

namespace App\Infrastructure\Project\Validation\Constraint;

use App\Domain\Workflow\Entity\Workflow;
use Override;
use Symfony\Component\Validator\Constraint;
use Symfony\Component\Validator\ConstraintValidator;
use Symfony\Component\Validator\Exception\UnexpectedTypeException;
use Symfony\Component\Validator\Exception\UnexpectedValueException;

final class MaxStepsPerWorkflowValidator extends ConstraintValidator
{
    #[Override]
    public function validate(mixed $value, Constraint $constraint): void
    {
        if (! $constraint instanceof MaxStepsPerWorkflow) {
            throw new UnexpectedTypeException($constraint, MaxStepsPerWorkflow::class);
        }

        if (null === $value) {
            return;
        }

        if (! $value instanceof Workflow) {
            throw new UnexpectedValueException($value, Workflow::class);
        }

        if ($value->getSteps()->count() >= MaxStepsPerWorkflow::MAX_STEPS) {
            $this->context->buildViolation($constraint->message)
                ->setCode(MaxStepsPerWorkflow::CODE)
                ->setParameter('{{ limit }}', (string) MaxStepsPerWorkflow::MAX_STEPS)
                ->addViolation();
        }
    }
}
