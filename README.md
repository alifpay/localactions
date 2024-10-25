# localactions
local GitHub Actions


# Act - Запуск GitHub Actions локально

[Act](https://nektosact.com) позволяет вам запускать ваши GitHub Actions локально.

## Установка

Для установки act выполните следующую команду:

```sh
curl --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```

## Использование

### Запуск всех workflow

Чтобы запустить все workflow, используйте команду:

```sh
act --pull=false
```

### Запуск конкретного workflow

Если у вас несколько workflow, вы можете сначала получить список всех доступных workflow:

```sh
act -l
```

### Запуск конкретного job

Чтобы запустить конкретный job из workflow, используйте флаг `-j` с указанием идентификатора job:

```sh
act -j <job_id> -P ubuntu-latest=node:20-bullseye --pull=false
```

Где `<job_id>` - это идентификатор job, который вы хотите запустить.

### Примеры

1. Запуск всех workflow:

    ```sh
    act --pull=false
    ```

2. Получение списка всех доступных workflow:

    ```sh
    act -l
    ```

3. Запуск конкретного job:

    ```sh
    act -j check -P ubuntu-latest=node:20-bullseye --pull=false
    ```
