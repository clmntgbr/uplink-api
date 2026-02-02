<?php

declare(strict_types=1);

namespace App\Infrastructure\Endpoint\Processor;

use ApiPlatform\Metadata\Operation;
use ApiPlatform\State\ProcessorInterface;
use App\Domain\Endpoint\Entity\Endpoint;
use App\Domain\User\Entity\User;
use Override;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\DependencyInjection\Attribute\Autowire;
use Symfony\Component\HttpKernel\Exception\BadRequestHttpException;
use Symfony\Component\HttpKernel\Exception\UnauthorizedHttpException;

/**
 * @implements ProcessorInterface<Endpoint, Endpoint>
 */
final readonly class CreateEndpointProcessor implements ProcessorInterface
{
    /**
     * @param ProcessorInterface<Endpoint, Endpoint> $persistProcessor
     */
    public function __construct(
        private readonly Security $security,
        #[Autowire(service: 'api_platform.doctrine.orm.state.persist_processor')]
        private ProcessorInterface $persistProcessor,
    ) {
    }

    #[Override]
    public function process(mixed $data, Operation $operation, array $uriVariables = [], array $context = []): Endpoint
    {
        $user = $this->security->getUser();

        if (! $user instanceof User) {
            throw new UnauthorizedHttpException('You have to be authenticated.');
        }

        if (! $data instanceof Endpoint) {
            throw new BadRequestHttpException('Invalid data.');
        }

        $data->setProject($user->getProject());

        return $this->persistProcessor->process($data, $operation, $uriVariables, $context);
    }
}
