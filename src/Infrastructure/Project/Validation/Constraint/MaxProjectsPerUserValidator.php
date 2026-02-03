<?php

declare(strict_types=1);

namespace App\Infrastructure\Project\Validation\Constraint;

use App\Domain\User\Entity\User;
use Override;
use Symfony\Component\Validator\Constraint;
use Symfony\Component\Validator\ConstraintValidator;
use Symfony\Component\Validator\Exception\UnexpectedTypeException;
use Symfony\Component\Validator\Exception\UnexpectedValueException;

final class MaxProjectsPerUserValidator extends ConstraintValidator
{
    #[Override]
    public function validate(mixed $value, Constraint $constraint): void
    {
        if (! $constraint instanceof MaxProjectsPerUser) {
            throw new UnexpectedTypeException($constraint, MaxProjectsPerUser::class);
        }

        if (null === $value) {
            return;
        }

        if (! $value instanceof User) {
            throw new UnexpectedValueException($value, User::class);
        }

        if ($value->getProjects()->count() >= MaxProjectsPerUser::MAX_PROJECTS) {
            $this->context->buildViolation($constraint->message)
                ->setParameter('{{ limit }}', (string) MaxProjectsPerUser::MAX_PROJECTS)
                ->addViolation();
        }
    }
}
