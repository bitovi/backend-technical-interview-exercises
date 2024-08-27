import Fastify from 'fastify';
import { PrismaClient } from '@prisma/client';

const fastify = Fastify({ logger: true });
const prisma = new PrismaClient();

interface OrderParams {
  id: string;
}

fastify.get('/orders', async (request, reply) => {
  try {
    const products = await prisma.product.findMany();
    const orders = await prisma.order.findMany();
    reply.send(orders.map((order) => ({
      ...order,
      products: order.products.map(
        (product) => products.find((p) => p.product_id === product)
    )
    })));
  } catch (error) {
    fastify.log.error(error);
    reply.status(500).send('Error retrieving orders');
  }
});

fastify.get<{ Params: OrderParams }>('/orders/:id', async (request, reply) => {
  const { id } = request.params;
  try {
    const order = await prisma.order.findUnique({
      where: { order_id: parseInt(id) },
    });

    if (!order) {
      reply.status(404).send('Order not found');
    } else {
      reply.send(order);
    }
  } catch (error) {
    fastify.log.error(error);
    reply.status(500).send('Error retrieving order');
  }
});

const start = async () => {
  try {
    await fastify.listen({ port: 3000, host: '0.0.0.0' });
    fastify.log.info('Server running at http://localhost:3000');
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
