<?php

declare(strict_types=1);

namespace App\Domain\Workflow\Doctrine;

use ApiPlatform\Doctrine\Orm\Extension\QueryCollectionExtensionInterface;
use ApiPlatform\Doctrine\Orm\Extension\QueryItemExtensionInterface;
use ApiPlatform\Doctrine\Orm\Util\QueryNameGeneratorInterface;
use ApiPlatform\Metadata\Operation;
use App\Domain\User\Entity\User;
use App\Domain\Workflow\Entity\Workflow;
use Doctrine\ORM\QueryBuilder;
use Exception;
use Override;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\HttpKernel\Exception\UnauthorizedHttpException;

use function sprintf;

final readonly class WorkflowExtension implements QueryCollectionExtensionInterface, QueryItemExtensionInterface
{
    public function __construct(
        private Security $security,
    ) {
    }

    /**
     * @throws Exception
     */
    #[Override]
    public function applyToCollection(QueryBuilder $queryBuilder, QueryNameGeneratorInterface $queryNameGenerator, string $resourceClass, ?Operation $operation = null, array $context = []): void
    {
        $this->addWhere($queryBuilder, $resourceClass);
    }

    /**
     * @throws Exception
     */
    #[Override]
    public function applyToItem(QueryBuilder $queryBuilder, QueryNameGeneratorInterface $queryNameGenerator, string $resourceClass, array $identifiers, ?Operation $operation = null, array $context = []): void
    {
        $this->addWhere($queryBuilder, $resourceClass);
    }

    /**
     * @throws Exception
     */
    private function addWhere(QueryBuilder $queryBuilder, string $resourceClass): void
    {
        if (Workflow::class !== $resourceClass) {
            return;
        }

        $user = $this->security->getUser();
        if (! $user instanceof User) {
            throw new UnauthorizedHttpException('You have to be authenticated.');
        }

        $rootAlias = $queryBuilder->getRootAliases()[0];
        $queryBuilder->andWhere(sprintf('%s.project = :project', $rootAlias));
        $queryBuilder->setParameter('project', $user->getProject());
    }
}
