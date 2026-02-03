<?php

declare(strict_types=1);

namespace App\Infrastructure\Workflow\Processor;

use ApiPlatform\Metadata\Operation;
use ApiPlatform\State\ProcessorInterface;
use App\Domain\User\Entity\User;
use App\Domain\Workflow\Entity\Workflow;
use Override;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\DependencyInjection\Attribute\Autowire;
use Symfony\Component\HttpKernel\Exception\BadRequestHttpException;
use Symfony\Component\HttpKernel\Exception\UnauthorizedHttpException;

/**
 * @implements ProcessorInterface<Workflow, Workflow>
 */
final readonly class CreateWorkflowProcessor implements ProcessorInterface
{
    /**
     * @param ProcessorInterface<Workflow, Workflow> $persistProcessor
     */
    public function __construct(
        private readonly Security $security,
        #[Autowire(service: 'api_platform.doctrine.orm.state.persist_processor')]
        private ProcessorInterface $persistProcessor,
    ) {
    }

    #[Override]
    public function process(mixed $data, Operation $operation, array $uriVariables = [], array $context = []): Workflow
    {
        $user = $this->security->getUser();

        if (! $user instanceof User) {
            throw new UnauthorizedHttpException('You have to be authenticated.');
        }

        if (! $data instanceof Workflow) {
            throw new BadRequestHttpException('Invalid data.');
        }

        $project = $user->getProject();
        $data->setProject($project);

        return $this->persistProcessor->process($data, $operation, $uriVariables, $context);
    }
}
