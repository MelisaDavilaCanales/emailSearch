Cada worker retorna el Resultado de la tarea realizada, incluye el resultado (value) pero tambien un error != nil en caso de existir.

Si existe canal de Result, Este resultado se envia por aqui, de modo que a este canal lo estara escuchando alguien mas, quien lo escuche tendra la oportunidad de evaluar si existe o no el error para en consecuencia de eso ejecutar o no la siguiente tarea.
=> Modularizacion, cada quien realiza su tarea segun lo que recibe.

=> Los workers solo escuchan, realizan tareas y mandan un resultado.

// PENDIENTE => Hacer que la processDirectory retorne un .Resuly[] para estandarizar.   
// Todas las funciones deben recibir un Result y evaluar si el err != nil

//Las funciones que se mandan a una WP que escucha de un Ch en el que escribio una WP debe recibir un Result y evaluar si el err != nil 

Y se va retornando sucesivamente el error inicial (entre funciones) de modo que si al final de la ejecucion se le informa al usuario, sepa el error inicial que hubo

// Batch con array de batchSize para eficiencia al agregar elementos, a diferencia del append

//WorkerPool y los tipos de Ch que debe recibir dependen del input y output de la TaskExecutor
// El tipo de lo que deben transportar tanto el canal que escuchan de tareas como en el/los que escriniran los resultados debe cumplir con el input de lo que la TaskExecutor de ese WP recibe (ya que se le envia lo que se escucha x el canal) asi como el resultado que retorna la taskExecutor porque sera lo que transporte el resulrCh

// Cada WP tiene su wg porque al terminar todos los workers de ese WP cierra sus canales correspondientes de Resultado x los que escribe