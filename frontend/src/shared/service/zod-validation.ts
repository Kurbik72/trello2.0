import * as z from 'zod'

const messageValidationSchema = {
  required: 'Fill in the blanks!',
  minLength: (count: number) => `Minimum number of characters ${count}`,
  maxLength: (count: number) => `Maximum number of characters ${count}`,
  email: 'Invalid email address',
  code: 'Enter code, please',
}

export const validationSchemas = z.object({
  required: z.string().min(1, {
    message: messageValidationSchema.required,
  }),
  minLength: (count: number) =>
    z.string().min(count, {
      message: messageValidationSchema.minLength(count),
    }),
  maxLength: (count: number) =>
    z.string().max(count, {
      message: messageValidationSchema.maxLength(count),
    }),
  email: z.string().email({
    message: messageValidationSchema.email,
  }),
  code: z.string().max(6, {
    message: messageValidationSchema.code,
  }),
})
export const validateService = <T>(data: T, schemaName: typeof validationSchemas) => {
  return schemaName.safeParse(data)
}
