<?php

declare(strict_types=1);

namespace App\Infrastructure\Step\Processor;

use ApiPlatform\Metadata\Operation;
use ApiPlatform\State\ProcessorInterface;
use App\Domain\Step\Entity\Step;
use App\Domain\User\Entity\User;
use Override;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\DependencyInjection\Attribute\Autowire;
use Symfony\Component\HttpKernel\Exception\BadRequestHttpException;
use Symfony\Component\HttpKernel\Exception\UnauthorizedHttpException;

/**
 * @implements ProcessorInterface<Step, Step>
 */
final readonly class CreateStepProcessor implements ProcessorInterface
{
    /**
     * @param ProcessorInterface<Step, Step> $persistProcessor
     */
    public function __construct(
        private readonly Security $security,
        #[Autowire(service: 'api_platform.doctrine.orm.state.persist_processor')]
        private ProcessorInterface $persistProcessor,
    ) {
    }

    #[Override]
    public function process(mixed $data, Operation $operation, array $uriVariables = [], array $context = []): Step
    {
        $user = $this->security->getUser();

        if (! $user instanceof User) {
            throw new UnauthorizedHttpException('You have to be authenticated.');
        }

        if (! $data instanceof Step) {
            throw new BadRequestHttpException('Invalid data.');
        }

        $workflow = $data->getWorkflow();

        $data->setPosition($workflow->getSteps()->count() + 1);

        return $this->persistProcessor->process($data, $operation, $uriVariables, $context);
    }
}
